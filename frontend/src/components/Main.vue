<script setup>
import {OpenFileDlgAndGenerateFile} from "../../wailsjs/go/main/App.js";

// 上传并生成文件
const uploadFile = () => {
  OpenFileDlgAndGenerateFile().then(()=>{console.log("ok")})
};

</script>

<template>
  <div class="file-upload">
    <div class="image-upload-wrap mt-1">
      <input class="file-upload-input" type='text' @click="uploadFile"  />
      <div class="drag-text">
        <h3>上传需要注入的ipa文件和Dylib文件</h3>
      </div>
    </div>
  </div>
</template>

<style scoped>

body {
  font-family: sans-serif;
  background-color: #eeeeee;
}

.file-upload {
  background-color: #ffffff;
  width: 300px;
  margin: 0 auto;
  padding: 20px;
}

.file-upload-input {
  position: absolute;
  margin: 0;
  padding: 0;
  width: 100%;
  height: 100%;
  outline: none;
  opacity: 0;
  cursor: pointer;
}

.image-upload-wrap {
  border: 4px dashed #1FB264;
  position: relative;
}

.image-upload-wrap:hover {
  background-color: #1FB264;
  border: 4px dashed #ffffff;
}


.drag-text {
  text-align: center;
}

.drag-text h3 {
  font-weight: 100;
  text-transform: uppercase;
  color: #15824B;
  padding: 60px 0;
}


</style>