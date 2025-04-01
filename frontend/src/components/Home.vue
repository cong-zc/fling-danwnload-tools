<script lang="ts" setup>
import { reactive } from "vue";
import { Greet, G_获取游戏名称 } from "../../wailsjs/go/main/App";

const data = reactive({
  搜索关键词: "",
  返回的数据: {
    标题: "",
    下载链接: "",
  },
});

function greet() {
  G_获取游戏名称(data.搜索关键词).then((result: any) => {
    data.返回的数据.标题 = result.标题;
    data.返回的数据.下载链接 = result.下载链接;
    console.log(result);
  });
}
</script>

<template>
  <main>
    <div id="input" class="input-box">
      <input
        id="name"
        v-model="data.搜索关键词"
        autocomplete="off"
        class="input"
        type="text"
      />
      <button class="btn" @click="greet">搜索</button>
    </div>
    <div id="result" class="result-box">
      <p>标题:{{ data.返回的数据.标题 }}</p>
      <p>
        下载链接:<a :href="data.返回的数据.下载链接">{{ data.返回的数据.下载链接 }}</a>
      </p>
    </div>
  </main>
</template>
