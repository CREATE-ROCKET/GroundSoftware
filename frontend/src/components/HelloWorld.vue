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
            <button class="btn" @click="getPortList()">Port Select</button>
          </div>
        </form>
        <form>
          <div class="input-section">
            <input type="text" v-model="dstId" class="input-field" placeholder="Destination ID">
            <input type="text" v-model="srcId" class="input-field" placeholder="Source ID">
          </div>
        </form>
        <form>
          <div class="input-section">
            <button @click="module_id($event)" class="btn">Set DstId and SrcId</button>
            <button @click="module_init($event)" class="btn">Send Module Command</button>
          </div>
        </form>
      </div>
    </div>
  </main>
  <footer>
    <div class="footer">
      <p>© 2023 <a href="https://github.com/CREATE-ROCKET">CREATE-ROCKET</a></p>
    </div>
  </footer>
</template>


<script setup>
import { reactive, onMounted, ref } from 'vue'
import { SerialStart } from '../../wailsjs/go/handler/App'
import { SerialStop } from '../../wailsjs/go/handler/App'
import { SerialSend } from '../../wailsjs/go/handler/App'
import { PortList } from '../../wailsjs/go/handler/App'
import { SelectedPort } from '../../wailsjs/go/handler/App'
import { ModuleStart } from '../../wailsjs/go/handler/App'
import { ModuleSend } from '../../wailsjs/go/handler/App'

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

function module_init(event) {
  event.preventDefault()
  if (sendMessage.value == '') {
    return
  }
  ModuleSend(sendMessage.value);
  if (isChecked.value) {
    clearMessage()
  }
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
    SerialSend(sendMessage.value)
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

let conn

function scrollToBottom() {
  const el = document.getElementById('monitor');
  var getBottomPosition = el.scrollHeight - el.scrollTop;
  if (getBottomPosition - 150 > el.clientHeight) {
    return
  }
  el.scrollTo(0, el.scrollHeight);
}

onMounted(() => {
  if (window.WebSocket) {
    conn = new WebSocket("ws://localhost:3007/ws")
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

<style scoped>
/* style.css */
.container {
  display: flex;
  height: 95svh;
  /* overflow-y: hidden; */
}

#monitor {
  border: 1px solid #ccc;
  padding: 20px;
  width: 50%;
  height: 100%;
  overflow-y: auto;
}

#inputs {
  padding: 20px;
  width: 50%;
  height: 100%;
  overflow-y: auto;
}

#message_table {
  width: 100%;
}

table {
  width: 100%;
  border-collapse: collapse;
}

td {
  border: 1px solid #ddd;
  padding: 8px;
}

.time {
  width: 20%;
}

.sender {
  width: 20%;
}

.content {
  width: 60%;
}

form {
  display: flex;
  flex-direction: column;
  margin-top: 5vh;
}

input[type=checkbox] {
  transform: scale(2);
  margin: 0 6px 0 0;
}

.input-section {
  margin-bottom: 10px;
}

.input-field {
  margin-bottom: 10px;
  padding: 8px;
  width: 100%;
  box-sizing: border-box;
}

.label {
  margin-bottom: 10px;
}

.reset {
  padding: 10px 20px;
  background-color: #ff675c;
  color: white;
  border: none;
  cursor: pointer;
  margin-right: 10px;
}

.reset:hover {
  background-color: #fe1100;
}

.btn {
  padding: 10px 20px;
  background-color: #4CAF50;
  color: white;
  border: none;
  cursor: pointer;
  margin-right: 10px;
}

.btn:hover {
  background-color: #45a049;
}

footer {
  position: absolute;
  bottom: 0;
  width: 50%;
  margin-left: 50%;
  height: 5vh;
  background-color: #f5f5f5;
  text-align: center;
}
</style>