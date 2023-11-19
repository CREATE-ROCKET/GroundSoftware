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
        <div class="status_field">
          <div class="status_left">
            <div class="status">
              <p>Volt Chart status: {{ loadVolt }}</p>
              <button @click="loadVolt_change($event)" class="btn">Volt Chart Change to {{ !loadVolt }}</button>
            </div>
            <div class="status">
              <p>Quat Chart status: {{ loadQuat }}</p>
              <button @click="loadQuat_change($event)" class="btn">Quat Chart Change to {{ !loadQuat }}</button>
            </div>
            <div class="status">
              <p>Lps Chart status: {{ loadLps }}</p>
              <button @click="loadLps_change($event)" class="btn">Lps Chart Change to {{ !loadLps }}</button>
            </div>
          </div>
          <div class="status_right">
            <div class="status">
              <p>OpenRate Chart status: {{ loadOpenRate }}</p>
              <button @click="loadOpenRate_change($event)" class="btn">OpenRate Chart Change to {{ !loadOpenRate
              }}</button>
            </div>
            <div class="status">
              <p>Model status: {{ loadModel }}</p>
              <button @click="loadModel_change($event)" class="btn">Model Change to {{ !loadModel }}</button>
            </div>
            <div class="status">
              <p>Quat Model Plot status: {{ quatStatus }}</p>
              <button @click="quatStatus_change($event)" class="btn">Quat Plot Change to {{ !quatStatus }}</button>
            </div>
          </div>
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
    <Cube :qua="quaternionArray" ref="cubeRef" />
  </main>
  <Footer />
</template>


<script>
import Cube from '/src/components/Quaternion.vue';
import ChartComponent from '/src/components/BarChart.vue';


export default {
  components: {
    Cube,
    ChartComponent,
  },
  data() {
    return {
      quaternionArray: [1, 0.5, 0, 0.5], // Example: Rotate by 90 degrees around the y-axis
    };
  },
  // methods: {
  //   rotateCube() {
  //     // Modify quaternionArray or call any other method in the Cube component
  //     // Example: rotate the cube by changing the quaternion values
  //     this.quaternionArray = [0, 0, 0.7071, 0.7071]; // Example: Rotate by 180 degrees around the z-axis
  //     // Access the Cube component using refs
  //     this.$refs.cubeRef.quaternion(); // Call the quaternion method in Cube
  //   },
  // },
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
import Footer from '/src/components/Footer.vue';

import Chart from '/src/components/BarChart.vue';

// import Quaternion from './Quaternion.vue'

// reactiveなデータプロパティを追加
const dstId = ref('');
const srcId = ref('');
const sendMessage = ref('');
const isChecked = ref(false);

function loadVolt_change(event) {
  event.preventDefault()
  if (loadVolt.value == true) {
    loadVolt.value = false
  } else {
    loadVolt.value = true
  }
}

function loadQuat_change(event) {
  event.preventDefault()
  if (loadQuat.value == true) {
    loadQuat.value = false
  } else {
    loadQuat.value = true
  }
}

function loadLps_change(event) {
  event.preventDefault()
  if (loadLps.value == true) {
    loadLps.value = false
  } else {
    loadLps.value = true
  }
}

function loadOpenRate_change(event) {
  event.preventDefault()
  if (loadOpenRate.value == true) {
    loadOpenRate.value = false
  } else {
    loadOpenRate.value = true
  }
}

function loadModel_change(event) {
  event.preventDefault()
  if (loadModel.value == true) {
    loadModel.value = false
  } else {
    loadModel.value = true
  }
}

function quatStatus_change(event) {
  event.preventDefault()
  if (quatStatus.value == true) {
    quatStatus.value = false
  } else {
    quatStatus.value = true
  }
}

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

const loadVolt = ref(false)
const loadQuat = ref(false)
const loadLps = ref(false)
const loadOpenRate = ref(false)

const loadModel = ref(false)

const quatStatus = ref(false)

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
let cubeRefComponent;

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
  function connectWebSocket() {
    if (window.WebSocket) {
      conn = new WebSocket("ws://localhost:3007/ws")
      chartComponentVoltage = getCurrentInstance().proxy.$refs.chartRefVoltage;
      chartComponentQuat1 = getCurrentInstance().proxy.$refs.chartRefQuat1;
      chartComponentQuat2 = getCurrentInstance().proxy.$refs.chartRefQuat2;
      chartComponentQuat3 = getCurrentInstance().proxy.$refs.chartRefQuat3;
      chartComponentQuat4 = getCurrentInstance().proxy.$refs.chartRefQuat4;
      chartComponentLps = getCurrentInstance().proxy.$refs.chartRefLps;
      chartComponentOpenRate = getCurrentInstance().proxy.$refs.chartRefOpenRate;
      cubeRefComponent = getCurrentInstance().proxy;

      conn.onclose = function (evt) {
        var time = new Date().toLocaleString()
        logMessages.push({ time, sender: "", content: "Connection closed." })

        // 再接続を試みる
        setTimeout(function () {
          connectWebSocket();
        }, 5000); // 5秒後に再接続試行
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
              if (loadVolt.value == true) {
                if (sender == "Voltage") {
                  voltage.value = content
                  let valuesArray = content.split(',');
                  chartComponentVoltage.addDataPoint(valuesArray[0], valuesArray[1]);
                }
              }
              if (loadQuat.value == true) {
                if (sender == "Quat") {
                  let valuesArray = content.split(',');
                  chartComponentQuat1.addDataPoint(valuesArray[0], valuesArray[1]);
                  chartComponentQuat2.addDataPoint(valuesArray[0], valuesArray[2]);
                  chartComponentQuat3.addDataPoint(valuesArray[0], valuesArray[3]);
                  chartComponentQuat4.addDataPoint(valuesArray[0], valuesArray[4]);
                }
              }
              if (loadLps.value == true) {
                if (sender == "Lps") {
                  let valuesArray = content.split(',');
                  chartComponentLps.addDataPoint(valuesArray[0], valuesArray[1]);
                }
              }
              if (loadOpenRate.value == true) {
                if (sender == "OpenRate") {
                  let valuesArray = content.split(',');
                  chartComponentOpenRate.addDataPoint(valuesArray[0], valuesArray[1]);
                }
              }
              if (loadModel.value == true) {
                if (sender == "Quat") {
                  let valuesArray = content.split(',');
                  cubeRefComponent.quaternionArray = [parseFloat(valuesArray[1]), parseFloat(valuesArray[2]), parseFloat(valuesArray[3]), parseFloat(valuesArray[4])];
                  cubeRefComponent.$refs.cubeRef.quaternion(quatStatus.value); // Call the quaternion method in Cube
                }
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
  }
  // WebSocket接続を開始
  connectWebSocket();
})
</script>

<style src="../style.css"></style>