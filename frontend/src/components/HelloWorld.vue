<template>
  <main>
    <div id="monitor" class="monitor">
      <div id="inputs" class="inputs">
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
      <div id="form">
        <form>
          <!-- <section id="buttons"> -->
            <button class="btn" @click="send($event)">Send</button>
            <input type="text" id="msg" size="64" autofocus />
            <br>
            <button class="btn" @click="serial_start($event)">Serial</button>
            <button class="btn" @click="serial_stop($event)">Serial Stop</button>
          <!-- </section> -->
        </form>
      </div> 
    </div>
  </main>
</template>

<script setup>
import { reactive, onMounted } from 'vue'
import { SerialStart } from '../../wailsjs/go/handler/App'
import { SerialStop } from '../../wailsjs/go/handler/App'
import { SerialSend } from '../../wailsjs/go/handler/App'

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
  var input = document.getElementById("msg")
  var message = input.value
  if (message.slice(0, 2) != "//") {
    SerialSend(message)
  }
  var timestamp = new Date().toLocaleString()
  if (message) {
    var messageWithClientId = `${timestamp}; ${clientId}:: ${message}`
    conn.send(messageWithClientId)
    input.value = ""
  }
}

let conn

function scrollToBottom() {
  const el = document.getElementById('message_table');
  var getBottomPosition  = el.scrollHeight - el.scrollTop;
  if (getBottomPosition -150 > el.clientHeight) {
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

table {
  width: auto;
  table-layout: fixed;
  word-wrap: break-word;
  border-collapse: collapse;
}

table,th {
  padding: 5px;
  text-align: left;
  /* min-width: 40vw; */
}

/* th.time {
  min-width: 30vw;
} */

/* th.sender {
  width: 10vw;
}

th.content {
  width: 60vw;
} */

#message_table {
  overflow: auto;
  height: 80vh;
  /* width: 100vw; */
}

/* #buttons {
  width: 30vw;
  bottom: 0;
  margin: 0;
} */

/* .result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
} */

.btn {
  /* width: 60px; */
  /* height: 30px; */
  line-height: 30px;
  border-radius: 30px;
  border: none;
  margin: 20 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.monitor {
  display: flex;
  flex-direction: column;
}

/* #log {
  height: auto;
  margin-bottom: 12px;
} */

#form {
  margin-top: 0;
  padding: 0;
  /* width: 30vw; */
  bottom: 0;
  height: 17vh;
  /* position: fixed; */
  /* display: -webkit-flex;
  display: -moz-flex;
  display: -ms-flex;
  display: -o-flex;
  display: flex; */
}

#inputs {
  /* position: fixed; */
  bottom: 0;
  margin: 0;
  height: 80vh;
  width: 33.3vw;
  display: -webkit-flex;
  display: -moz-flex;
  display: -ms-flex;
  display: -o-flex;
  display: flex;
  margin-bottom: 12px;
}

.btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>