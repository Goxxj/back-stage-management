<template>
  <div class="tabs">
    <el-tag
      v-for="(item, index) in tags"
      :key="item.label"
      :closable="item.name !== 'home'"
      @click="changeMenu(item)"
      @close="handerClose(item, index)"
      size="small"
      :effect="$route.name === item.name ? 'dark' : 'plain'"
    >
      {{ item.label }}
    </el-tag>
  </div>
</template>

<script>
import { mapState } from "vuex";
export default {
  name: "CommonTag",
  data() {
    return {};
  },
  computed: {
    ...mapState({
      tags: (state) => state.tab.tabsList,
    }),
  },
  methods: {
    //点击TAG跳转
    changeMenu(item) {
      this.$router.push({
        name: item.name,
      });
    },
    //点击Tag删除
    handerClose(item, index) {
      this.$store.commit("closeTag", item);
      if (item.name !== this.$route.name) {
        return;
      }
      if (index === this.tags.length) {
        this.$router.push({
          name: this.tags[index - 1].name,
        });
      } else {
        this.$router.push({
          name: this.tags[index].name,
        });
      }
    },
  },
};
</script>

<style lang="less" scoped>
.tabs {
  padding: 20px;
  background-color: #e9eef3;
  .el-tag {
    margin-right: 15px;
    cursor: pointer;
  }
}
</style>