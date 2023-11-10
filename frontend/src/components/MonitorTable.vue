<template>
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
</template>

<script>
import { ref } from 'vue';

export default {
  setup() {
    const logMessages = ref([]);

    return {
      logMessages,
    };
  },
};
</script>