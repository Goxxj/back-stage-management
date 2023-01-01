<template>
  <el-card class="login"
    ><el-form
      ref="register"
      :model="form"
      :rules="rules"
      :inline="true"
      label-width="70px"
    >
      <h3>系统注册</h3>
      <el-form-item label="用户名" prop="username">
        <el-input v-model="form.username" placeholder="请输入用户名"></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input
          type="password"
          v-model="form.password"
          placeholder="请输入密码"
        ></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submit">提交</el-button>
      </el-form-item>
    </el-form>
  </el-card>
</template>
  
  <script>
import { registerUser } from "../api/index";
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Register",
  data() {
    return {
      form: {
        username: "",
        password: "",
      },
      rules: {
        username: [
          { required: true, message: "请输入用户名", trigger: "blur" },
          { min: 3, max: 5, message: "长度在 3 到 5 个字符", trigger: "blur" },
        ],
        password: [{ required: true, message: "请输入密码", trigger: "blur" }],
      },
    };
  },
  methods: {
    submit() {
      this.$refs.register.validate((valid) => {
        if (valid) {
          registerUser(this.form).then((data) => {
            console.log(data, "data");
            if (data.data.result === "注册成功") {
              console.log(1);
              this.$message(data.data.result);
              setTimeout(() => {
                this.$router.push("login");
              }, 1000);
            } else {
              console.log(2);
              this.$message.error(data.data.result);
            }
          });
        }
      });
    },
  },
};
</script>
  
  <style lang="less" scoped>
.login {
  position: relative;
  box-sizing: border-box;
  width: 350px;
  margin: 180px auto;
  background-color: #fff;
  padding: 35px 35px 15px 35px;
  border-radius: 15px;

  /deep/.el-card__body,
  .el-main {
    padding: 0;
  }

  h3 {
    margin: 3px;
    color: #505458;
    text-align: center;
  }
  .el-form-item {
    .el-input {
      width: 198px;
    }
    .el-button {
      margin-left: 105px;
    }
  }
  .register {
    position: absolute;
    top: 203px;
    left: 247px;
    cursor: pointer;
  }
}
</style>