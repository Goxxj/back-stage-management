<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <div class="user">
    <el-dialog title="提示" :visible.sync="dialogVisible" width="50%">
      <el-form
        ref="form"
        :inline="true"
        :model="form"
        label-width="100px"
        size="mini"
        :rules="rules"
        :before-close="handleClose"
        style="text-align: left"
      >
        <el-form-item label="姓名" prop="name">
          <el-input v-model="form.name" placeholder="请输入姓名"></el-input>
        </el-form-item>
        <el-form-item label="年龄" prop="age">
          <el-input v-model="form.age" placeholder="请输入年龄"></el-input>
        </el-form-item>
        <el-form-item label="性别" prop="sex">
          <el-select v-model="form.sex" placeholder="请选择">
            <el-option label="男" value="男"></el-option>
            <el-option label="女" value="女"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="出生日期" prop="birth">
          <el-date-picker
            v-model="form.birth"
            type="date"
            placeholder="选择日期"
            value-format="yyyy-MM-dd"
          >
          </el-date-picker>
        </el-form-item>
        <el-form-item label="地址" prop="addr">
          <el-input v-model="form.addr" placeholder="请输入地址"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="cancel">取 消</el-button>
        <el-button type="primary" @click="submit">确 定</el-button>
      </span>
    </el-dialog>
    <div class="user-header">
      <el-button @click="handleAdd" type="primary" id="btn"> +新增 </el-button>
      <!-- form搜索区 -->
      <el-form :inline="true" :model="userForm">
        <el-form-item>
          <el-input v-model="userForm.name" placeholder="请输入名称"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">查询</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="user-tab">
      <el-table :data="tablePagination" stripe style="width: 100%">
        <el-table-column prop="name" label="姓名"> </el-table-column>
        <el-table-column prop="sex" label="性别"> </el-table-column>
        <el-table-column prop="age" label="年龄"> </el-table-column>
        <el-table-column prop="birth" label="出生日期"> </el-table-column>
        <el-table-column prop="addr" label="地址"> </el-table-column>
        <el-table-column prop="addr" label="操作">
          <template slot-scope="scope">
            <el-button size="mini" @click="handleEdit(scope.row)"
              >编辑</el-button
            >
            <el-button
              size="mini"
              type="danger"
              @click="handleDelete(scope.row)"
              >删除</el-button
            >
          </template>
        </el-table-column>
      </el-table>
      <div class="fen">
        <el-pagination
          layout="prev, pager, next"
          :total="total"
          hide-on-single-page
          @current-change="handlePage"
        >
        </el-pagination>
      </div>
    </div>
  </div>
</template>

<script>
import { getUser, addUser, editUser, deleteUser } from "@/api";
export default {
  data() {
    return {
      dialogVisible: false,
      form: {
        name: "",
        age: "",
        sex: "",
        birth: "",
        addr: "",
      },
      rules: {
        name: [{ required: true, message: "请输入名字" }],
        age: [{ required: true, message: "请输入年龄" }],
        sex: [{ required: true, message: "请输入性别" }],
        birth: [{ required: true, message: "请输入出生日期" }],
        addr: [{ required: true, message: "请输入地址" }],
      },
      tableData: [],
      tablePagination: [],
      moduleType: 0, //0表示新增，1表示编辑
      total: 0,
      userForm: {
        name: "",
      },
    };
  },
  methods: {
    submit() {
      //校验
      this.$refs.form.validate((valid) => {
        if (valid) {
          //新增
          if (this.moduleType === 0) {
            addUser(this.form).then((data) => {
              this.$message(data.data.result);
              this.getList();
            });
          } else {
            //编辑行
            editUser(this.form).then((data) => {
              this.$message(data.data.result);
              this.getList();
            });
          }
          //清空表单
          this.$refs.form.resetFields();
          //关闭表单
          this.dialogVisible = false;
        }
      });
    },
    //弹窗关闭时
    handleClose() {
      this.$refs.form.resetFields();
      this.dialogVisible = false;
    },
    cancel() {
      this.handleClose();
    },
    //编辑行
    handleEdit(row) {
      this.moduleType = 1;
      this.dialogVisible = true;
      this.form = JSON.parse(JSON.stringify(row));
    },
    //删除行
    handleDelete(row) {
      this.$confirm("此操作将永久删除该文件, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          deleteUser(row).then((data) => {
            this.$message({
              type: "success",
              message: data.data.result,
            });
            this.getList();
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除",
          });
        });
    },
    //新增用户
    handleAdd() {
      this.moduleType = 0;
      this.dialogVisible = true;
    },
    //获取列表
    getList(data) {
      getUser(data).then((data) => {
        this.tableData = data.data;
        if (this.tableData.length > 10) {
          this.tablePagination = this.tableData.slice(0, 10);
        } else {
          this.tablePagination = this.tableData;
        }

        this.total = data.data ? data.data.length : 0;
      });
    },
    //分页切片
    handlePage(page) {
      if (this.tableData.length > page * 10) {
        this.tablePagination = this.tableData.slice(10 * (page - 1), page * 10);
      } else {
        this.tablePagination = this.tableData.slice(10 * (page - 1));
      }
    },
    onSubmit() {
      this.getList(this.userForm);
    },
  },

  mounted() {
    //获取用户信息
    this.getList(this.userForm);
  },
};
</script>

<style lang="less" scoped>
.user {
  height: 90%;
  .user-tab {
    position: relative;

    .fen {
      position: absolute;
      bottom: -25px;
      right: 0;
    }
  }
}
.user-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  #btn {
    width: 78.18px;
  }
}
</style>