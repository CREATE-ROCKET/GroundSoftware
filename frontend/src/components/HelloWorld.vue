<template>
  <main>
    <div class="container">
      <div id="monitor" class="monitor">
        <div id="message_table">
          <table>
            <tbody>
              <tr v-for="(message, index) in logMessages" :key="index">
                <td class="time">{{ message.time }}</td>
                <td class="sender">{{ message.sender }}</td>
                <td class="content">{{ message.content }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
      <div id="inputs" class="inputs">
        <form>
          <div class="input-section">
            <label for="msg" class="label">Message:</label>
            <input type="text" v-model="sendMessage" id="msg" class="input-field" autofocus />
            <!-- <input type="reset" value="Reset" id="msg" class="reset" autofocus /> -->
            <input type="submit" id="msg" class="btn" @click="send($event)" value='Send' />
            <label for="switch">
              <input type="checkbox" id="switch" v-model="isChecked" />
              Message Clear
            </label>
          </div>
        </form>
        <form>
          <div class="input-section">
            <button class="btn" @click="serial_start($event)">Serial Start</button>
            <button class="btn" @click="serial_stop($event)">Serial Stop</button>
          </div>
        </form>
        <form>
          <div class="input-section">
            <label for="port" class="label">Port:</label>
            <select v-model="state.selected" @change="selected_port(state.selected)">
              <option v-for="port in state.portList" :key="port">{{ port }}</option>
            </select>
            <button class="btn" @click="getPortList()">Get Port List</button>
          </div>
        </form>
        <form>
          <div class="input-section">
            <input type="text" v-model="dstId" class="input-field" placeholder="Dst ID">
            <input type="text" v-model="srcId" class="input-field" placeholder="Src ID">
          </div>
        </form>
        <form>
          <div class="input-section">
            <button @click="module_id($event)" class="btn">Set DstId and SrcId</button>
            <button @click="module_send($event)" class="btn">Send Module Command</button>
            <button @click="module_config($event)" class="btn">Set Module From Config File</button>
          </div>
        </form>
        <div>
          <p>Voltage: {{ voltage }}</p>
        </div>
      </div>
    </div>
    <div class="chart-section">
      <p>Voltage</p>
      <Chart ref="chartRefVoltage" />
      <p>Quat1</p>
      <Chart ref="chartRefQuat1" />
      <p>Quat2</p>
      <Chart ref="chartRefQuat2" />
    </div>
    <div class="chart-section">
      <p>Quat3</p>
      <Chart ref="chartRefQuat3" />
      <p>Quat4</p>
      <Chart ref="chartRefQuat4" />
      <p>LPS</p>
      <Chart ref="chartRefLps" />
    </div>
    <div class="chart-section">
      <p>OpenRate</p>
      <Chart ref="chartRefOpenRate" />
    </div>
  </main>
  <Footer />
</template>


<script>
import ChartComponent from './BarChart.vue';

export default {
  components: {
    ChartComponent,
  },
  methods: {
  },
};
</script>



<script setup>
import { reactive, onMounted, ref, getCurrentInstance } from 'vue'
import { SerialStart } from '../../wailsjs/go/handler/App'
import { SerialStop } from '../../wailsjs/go/handler/App'
import { SerialTextSend } from '../../wailsjs/go/handler/App'
import { PortList } from '../../wailsjs/go/handler/App'
import { SelectedPort } from '../../wailsjs/go/handler/App'
import { ModuleStart } from '../../wailsjs/go/handler/App'
import { ModuleSend } from '../../wailsjs/go/handler/App'
import { ModuleEnv } from '../../wailsjs/go/handler/App'
import Footer from './Footer.vue';

import Chart from './BarChart.vue';

// reactiveなデータプロパティを追加
const dstId = ref('');
const srcId = ref('');
const sendMessage = ref('');
const isChecked = ref(false);

// module_start関数を更新し、データプロパティから引数を使用する
function module_id(event) {
  event.preventDefault()
  ModuleStart(dstId.value, srcId.value);
}

function module_send(event) {
  event.preventDefault()
  if (sendMessage.value == '') {
    return
  }
  ModuleSend(sendMessage.value);
  if (isChecked.value) {
    clearMessage()
  }
}

function module_config(event) {
  event.preventDefault()
  ModuleEnv();
}

const state = reactive({
  portList: [],
  selected: null,
  error: null,
});

onMounted(async () => {
  try {
    const ports = await getPortList();
    state.portList = ports;
  } catch (error) {
    state.error = error;
  }
});

async function getPortList() {
  try {
    const portList = await PortList();
    return portList;
  } catch (error) {
    throw new Error('Error fetching port list: ' + error);
  }
}

const selected_port = async (port) => {
  try {
    await SelectedPort(port);
  } catch (error) {
    state.error = 'Error selecting port: ' + error;
  }
};

function serial_start(event) {
  event.preventDefault()
  SerialStart()
}

function serial_stop(event) {
  event.preventDefault()
  SerialStop()
  scrollToBottom()
}

const logMessages = reactive([])
const voltage = ref(0)
let clientId = "User"

function send(event) {
  event.preventDefault()
  if (sendMessage.value == '') {
    return
  }
  // var input = document.getElementById("msg")
  // var message = input.value
  // var message = sendMessage.value
  if (sendMessage.value.slice(0, 2) != "//") {
    SerialTextSend(sendMessage.value)
  }
  var timestamp = new Date().toLocaleString()
  if (sendMessage.value) {
    var messageWithClientId = `${timestamp}; ${clientId}:: ${sendMessage.value}`
    conn.send(messageWithClientId)
  }
  if (isChecked.value) {
    clearMessage()
  }
}

function clearMessage() {
  sendMessage.value = ""
}

let conn;
let chartComponentVoltage;
let chartComponentQuat1;
let chartComponentQuat2;
let chartComponentQuat3;
let chartComponentQuat4;
let chartComponentLps;
let chartComponentOpenRate;

function scrollToBottom() {
  const el = document.getElementById('monitor');
  var getBottomPosition = el.scrollHeight - el.scrollTop;
  if (getBottomPosition - 150 > el.clientHeight) {
    return
  }
  el.scrollTo(0, el.scrollHeight);
}

onMounted(() => {
  // const chartComponent = this.$refs.chartRef;
  if (window.WebSocket) {
    conn = new WebSocket("ws://localhost:3007/ws")
    chartComponentVoltage = getCurrentInstance().proxy.$refs.chartRefVoltage;
    chartComponentQuat1 = getCurrentInstance().proxy.$refs.chartRefQuat1;
    chartComponentQuat2 = getCurrentInstance().proxy.$refs.chartRefQuat2;
    chartComponentQuat3 = getCurrentInstance().proxy.$refs.chartRefQuat3;
    chartComponentQuat4 = getCurrentInstance().proxy.$refs.chartRefQuat4;
    chartComponentLps = getCurrentInstance().proxy.$refs.chartRefLps;
    chartComponentOpenRate = getCurrentInstance().proxy.$refs.chartRefOpenRate;

    conn.onclose = function (evt) {
      var time = new Date().toLocaleString()
      logMessages.push({ time, sender: "", content: "Connection closed." })
    }

    conn.onmessage = function (evt) {
      const messages = evt.data.split('\n')
      for (let i = 0; i < messages.length; i++) {
        const message = messages[i]
        if (message == "") {
          continue
        }
        const separatorIndex = message.indexOf(';') // 時刻とメッセージの区切り位置を探す
        if (separatorIndex !== -1) {
          var time = message.substring(0, separatorIndex) // 時刻を抽出
          const content = message.substring(separatorIndex + 1) // メッセージ本文を抽出

          const separatorIndex2 = message.indexOf('::') // 識別子とメッセージの区切り位置を探す
          if (separatorIndex2 !== -1) {
            const sender = message.substring(separatorIndex + 1, separatorIndex2) // 識別子を抽出
            const content = message.substring(separatorIndex2 + 1 + 1) // メッセージ本文を抽出
            logMessages.push({ time, sender, content }) // 送信者とメッセージを表示
          } else {
            logMessages.push({ time, sender: "Anonymous", content })
          }
        } else {
          const separatorIndex2 = message.indexOf('::') // 識別子とメッセージの区切り位置を探す
          if (separatorIndex2 !== -1) {
            const sender = message.substring(0, separatorIndex2) // 識別子を抽出
            const content = message.substring(separatorIndex2 + 1 + 1) // メッセージ本文を抽出
            var time = new Date().toLocaleString()
            logMessages.push({ time, sender, content }) // 送信者とメッセージを表示
            if (sender == "Voltage") {
              voltage.value = content
              let valuesArray = content.split(',');
              chartComponentVoltage.addDataPoint(valuesArray[0], valuesArray[1]);
            }
            if (sender == "Quat") {
              let valuesArray = content.split(',');
              chartComponentQuat1.addDataPoint(valuesArray[0], valuesArray[1]);
              chartComponentQuat2.addDataPoint(valuesArray[0], valuesArray[2]);
              chartComponentQuat3.addDataPoint(valuesArray[0], valuesArray[3]);
              chartComponentQuat4.addDataPoint(valuesArray[0], valuesArray[4]);
            }
            if (sender == "Lps") {
              let valuesArray = content.split(',');
              chartComponentLps.addDataPoint(valuesArray[0], valuesArray[1]);
            }
            if (sender == "OpenRate") {
              let valuesArray = content.split(',');
              chartComponentOpenRate.addDataPoint(valuesArray[0], valuesArray[1]);
            }
          } else {
            var time = new Date().toLocaleString()
            logMessages.push({ time, sender: "", content: message })
          }
        }
        scrollToBottom()
      }
    }
  } else {
    var time = new Date().toLocaleString()
    logMessages.push({ time, sender: "", content: "Your browser does not support WebSockets." })
  }
})
</script>


<style src="../style.css"></style>