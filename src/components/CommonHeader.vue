<template>
  <div class="header-container">
    <div class="l-content">
      <el-button
        icon="el-icon-menu"
        @click="handerMenu"
        size="mini"
      ></el-button>

      <el-breadcrumb separator="/">
        <el-breadcrumb-item
          v-for="item in tags"
          :key="item.name"
          :to="{ path: item.path }"
          >{{ item.label }}</el-breadcrumb-item
        >
      </el-breadcrumb>
    </div>
    <div class="r-content">
      <el-dropdown @command="handleClick">
        <span class="el-dropdown-link">
          <img src="../assets/xiaoyu.png" />
        </span>
        <el-dropdown-menu slot="dropdown">
          <el-dropdown-item>个人中心</el-dropdown-item>
          <el-dropdown-item command="cancel">退出</el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";
import Cookie from "js-cookie";
export default {
  data() {
    return {};
  },
  methods: {
    handerMenu() {
      this.$store.commit("collapseMenu");
    },
    handleClick(command) {
      if (command === "cancel") {
        //清除cookie
        Cookie.remove("token");
        Cookie.remove("menu");
        this.$router.push("login");
      }
    },
  },
  computed: {
    ...mapState({
      tags: (state) => state.tab.tabsList,
    }),
  },
};
</script>

<style lang="less" scoped>
.header-container {
  padding: 0 20px;
  background-color: #333;
  height: 60px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  .el-button {
    margin-right: 20px;
  }
  .shouye {
    padding: 20px;
    color: #fff;
    font-size: 14px;
  }
  .l-content {
    display: flex;
    align-items: center;
    /deep/.el-breadcrumb__item {
      .el-breadcrumb__inner {
        color: #666;
      }
      &:last-child {
        .el-breadcrumb__inner {
          color: #fff;
        }
      }
    }
  }
}
.el-dropdown-link {
  padding: 0;
  cursor: pointer;
  color: #409eff;
  img {
    margin-top: 20px;
    width: 40px;
    border-radius: 50%;
  }
}

.el-icon-arrow-down {
  font-size: 12px;
}
</style>