<template>
  <main>
    <div id="monitor" class="monitor">
      <dev id="log" class="log">
        <div v-for="(message, index) in logMessages" :key="index">{{ message }}</div>
      </dev>
      <dev id="inputs" class="inputs">
        <button class="btn" @click="serial($event)">Serial</button>
        <button class="btn" @click="serial_stop($event)">Serial Stop</button>
        <form id="form">
          <button class="btn" @click="send($event)">Send</button>
          <input type="text" id="msg" size="64" autofocus />
        </form>
      </dev> 
    </div>
  </main>
</template>

<script setup>
import { reactive, onMounted } from 'vue'
import { SerialMonitor } from '../../wailsjs/go/main/App'
import { SerialStop } from '../../wailsjs/go/main/App'

function serial(event) {
  event.preventDefault()
  SerialMonitor()
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
  var timestamp = new Date().toLocaleString()
  if (message) {
    var messageWithClientId = `${timestamp}; ${clientId}:: ${message}`
    conn.send(messageWithClientId)
    input.value = ""
  }
}

let conn

function scrollToBottom() {
  // document.getElementById("log").scrollIntoView(false)
  // document.getElementById("monitor").scrollIntoView(false)
  // document.getElementById("inputs").scrollIntoView(false)
  document.scrollIntoView(false)
}

onMounted(() => {
  if (window.WebSocket) {
    conn = new WebSocket("ws://localhost:3007/ws")
    conn.onclose = function (evt) {
      logMessages.push("Connection closed.")
    }

    conn.onmessage = function (evt) {
      const messages = evt.data.split('\n')
      for (let i = 0; i < messages.length; i++) {
        const message = messages[i]
        const separatorIndex = message.indexOf(';') // 時刻とメッセージの区切り位置を探す
        if (separatorIndex !== -1) {
          const time = message.substring(0, separatorIndex) // 時刻を抽出
          const content = message.substring(separatorIndex + 1) // メッセージ本文を抽出

          const separatorIndex2 = message.indexOf('::') // 識別子とメッセージの区切り位置を探す
          if (separatorIndex2 !== -1) {
            const sender = message.substring(separatorIndex + 1, separatorIndex2) // 識別子を抽出
            const content = message.substring(separatorIndex2 + 1) // メッセージ本文を抽出
            logMessages.push(`${time} From ${sender} ${content}`) // 送信者とメッセージを表示
          } else {
            logMessages.push(`${time} From Anonymous: ${content}`)
          }
        } else {
          const separatorIndex2 = message.indexOf('::') // 識別子とメッセージの区切り位置を探す
          if (separatorIndex2 !== -1) {
            const sender = message.substring(0, separatorIndex2) // 識別子を抽出
            const content = message.substring(separatorIndex2 + 1) // メッセージ本文を抽出
            const time = new Date().toLocaleString()
            logMessages.push(`${time} From ${sender} ${content}`) // 送信者とメッセージを表示
          } else {
            logMessages.push(message)
          }
        }
      scrollToBottom()
      }
    }
  } else {
    logMessages.push("Your browser does not support WebSockets.")
  }
})
</script>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.monitor {
  display: flex;
  flex-direction: column;
}

#log {
  height: auto;
  margin-bottom: 8%;
}

#inputs {
  position: fixed;
  bottom: 0;
  margin: 0;
  height: auto;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>