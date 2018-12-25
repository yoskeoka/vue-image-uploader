<template>
  <div class="hello">
    <vue-dropzone
      ref="myVueDropzone"
      id="dropzone"
      :options="dropzoneOptions"
      v-on:vdropzone-sending="sendingEvent"
      v-on:vdropzone-removed-file="removeEvent"
    ></vue-dropzone>
  </div>
</template>

<script>
import vue2Dropzone from "vue2-dropzone";
import axios from "axios";
import "vue2-dropzone/dist/vue2Dropzone.min.css";
export default {
  name: "HelloWorld",
  data: function() {
    return {
      dropzoneOptions: {
        url: `http://localhost:8888/images`,
        method: "post",
        addRemoveLinks: true
      }
    };
  },
  components: {
    vueDropzone: vue2Dropzone
  },
  methods: {
    sendingEvent: function(file, xhr, formData) {
      formData.append("uuid", file.upload.uuid);
    },
    removeEvent: function(file, error, xhr) {
      axios
        .delete(`http://localhost:8888/images/${file.upload.uuid}`)
        .then(res => {
          console.log(res.data);
        })
        .catch(err => {
          console.error(err);
        });
    }
  },
  mounted(){
    axios.get("http://localhost:8888/images").then(res => {
      res.data.forEach(res => {
        const filename = res.path.split("/").pop();
        const ext = filename.split(".").pop();
        let uuid = filename === ext ? filename : filename.replace("."+ext, "");
        const file = {size: res.size, name: filename, type: `image/${ext}`, upload: {uuid}}
        this.$refs.myVueDropzone.manuallyAddFile(file, res.path);
      })
    }).catch(err => {
      console.error(err);
    })
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
