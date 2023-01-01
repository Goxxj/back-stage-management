<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <el-row>
    <el-col :span="8" style="padding-right: 10px">
      <el-card class="box-card">
        <div class="user">
          <img src="../assets/xiaoyu.png" alt="" />
          <div class="userInfo">
            <p class="name">Admin</p>
            <p class="access">超级管理员</p>
          </div>
        </div>
        <div class="login-info">
          <p>
            上次登陆的时间：
            <span>2022年11月8日</span>
          </p>
          <p>
            上次登录的地点：
            <span>成都</span>
          </p>
        </div>
      </el-card>
      <el-card style="margin-top: 20px; height: 460px">
        <el-table :data="tableData" style="width: 100%">
          <el-table-column
            v-for="(val, key) in tableLable"
            :key="key"
            :prop="key"
            :label="val"
          >
          </el-table-column>
        </el-table>
      </el-card>
    </el-col>
    <el-col :span="16" style="padding-left: 10px">
      <div class="num">
        <el-card
          :body-style="{ display: 'flex ', padding: 0 }"
          v-for="item in countData"
          :key="item.name"
        >
          <i
            class="icon"
            :class="`el-icon-${item.icon}`"
            :style="{ background: item.color }"
          ></i>
          <div class="detail">
            <p class="num">${{ item.value }}</p>
            <p class="str">{{ item.name }}</p>
          </div>
        </el-card>
      </div>
      <div>
        <el-card style="height: 280px">
          <div ref="echarts1" style="height: 280px"></div>
        </el-card>
        <div class="graph">
          <el-card style="height: 260px">
            <dir ref="echarts2" style="height: 260px"></dir>
          </el-card>
          <el-card style="height: 260px">
            <div ref="echarts3" style="height: 240px"></div>
          </el-card>
        </div>
      </div>
    </el-col>
  </el-row>
</template>

<script>
import { getData } from "../api/index";
import * as echarts from "echarts";
import randomRange from "../math/index";
export default {
  data() {
    return {
      tableData: [],
      tableLable: {},
      countData: [
        {
          name: "今日支付订单",
          value: 1234,
          icon: "success",
          color: "#2ec7c9",
        },
        {
          name: "今日收藏订单",
          value: 210,
          icon: "star-on",
          color: "#ffb980",
        },
        {
          name: "今日未支付订单",
          value: 1234,
          icon: "s-goods",
          color: "#5ab1ef",
        },
        {
          name: "本月支付订单",
          value: 1234,
          icon: "success",
          color: "#2ec7c9",
        },
        {
          name: "本月收藏订单",
          value: 210,
          icon: "star-on",
          color: "#ffb980",
        },
        {
          name: "本月未支付订单",
          value: 1234,
          icon: "s-goods",
          color: "#5ab1ef",
        },
      ],
      videoData: [
        {
          name: "小米",
          value: 2999,
        },
        {
          name: "苹果",
          value: 5999,
        },
        {
          name: "vivo",
          value: 1500,
        },
        {
          name: "oppo",
          value: 1999,
        },
        {
          name: "魅族",
          value: 2200,
        },
        {
          name: "三星",
          value: 4500,
        },
      ],
      userData: [
        {
          date: "周一",
          new: 5,
          active: 200,
        },
        {
          date: "周二",
          new: 10,
          active: 500,
        },
        {
          date: "周三",
          new: 12,
          active: 550,
        },
        {
          date: "周四",
          new: 60,
          active: 800,
        },
        {
          date: "周五",
          new: 65,
          active: 550,
        },
        {
          date: "周六",
          new: 53,
          active: 770,
        },
        {
          date: "周日",
          new: 33,
          active: 170,
        },
      ],
      orderData: {
        date: [
          "20191001",
          "20191002",
          "20191003",
          "20191004",
          "20191005",
          "20191006",
          "20191007",
        ],
        data: [
          {
            苹果: randomRange(0, 10000),
            vivo: randomRange(0, 10000),
            oppo: randomRange(0, 10000),
            魅族: randomRange(0, 10000),
            三星: randomRange(0, 10000),
            小米: randomRange(0, 10000),
          },
          {
            苹果: randomRange(0, 10000),
            vivo: randomRange(0, 10000),
            oppo: randomRange(0, 10000),
            魅族: randomRange(0, 10000),
            三星: randomRange(0, 10000),
            小米: randomRange(0, 10000),
          },
          {
            苹果: randomRange(0, 10000),
            vivo: randomRange(0, 10000),
            oppo: randomRange(0, 10000),
            魅族: randomRange(0, 10000),
            三星: randomRange(0, 10000),
            小米: randomRange(0, 10000),
          },
          {
            苹果: randomRange(0, 10000),
            vivo: randomRange(0, 10000),
            oppo: randomRange(0, 10000),
            魅族: randomRange(0, 10000),
            三星: randomRange(0, 10000),
            小米: randomRange(0, 10000),
          },
          {
            苹果: randomRange(0, 10000),
            vivo: randomRange(0, 10000),
            oppo: randomRange(0, 10000),
            魅族: randomRange(0, 10000),
            三星: randomRange(0, 10000),
            小米: randomRange(0, 10000),
          },
          {
            苹果: randomRange(0, 10000),
            vivo: randomRange(0, 10000),
            oppo: randomRange(0, 10000),
            魅族: randomRange(0, 10000),
            三星: randomRange(0, 10000),
            小米: randomRange(0, 10000),
          },
          {
            苹果: randomRange(0, 10000),
            vivo: randomRange(0, 10000),
            oppo: randomRange(0, 10000),
            魅族: randomRange(0, 10000),
            三星: randomRange(0, 10000),
            小米: randomRange(0, 10000),
          },
        ],
      },
    };
  },
  beforeMount() {
    getData().then((data) => {
      this.tableData = data.data.tableData;
      this.tableLable = data.data.tableLable;
      this.countData = data.data.countData;
      //图表数据
    });
  },
  mounted() {
    //折线图
    const echarts1 = echarts.init(this.$refs.echarts1);
    var echarts1Option = {};
    const orderData = this.orderData;

    const xAxis = Object.keys(orderData.data[0]);
    const xAxisData = {
      data: xAxis,
    };
    //X轴
    echarts1Option.xAxis = {
      data: orderData.date,
    };
    //Y
    echarts1Option.yAxis = {};
    echarts1Option.legend = xAxisData;
    echarts1Option.series = [];
    xAxis.forEach((key) => {
      echarts1Option.series.push({
        name: key,
        data: orderData.data.map((item) => item[key]),
        type: "line",
      });
    });
    echarts1.setOption(echarts1Option);

    //柱状图
    const echarts2 = echarts.init(this.$refs.echarts2);
    const echarts2Option = {
      lengend: {
        textStyle: {
          color: "#333",
        },
      },
      grid: {
        left: "20%",
      },
      tooltip: {
        trigger: "axis",
      },
      xAxis: {
        type: "category",
        data: this.userData.map((item) => item.date),
        axisLine: {
          lineStyle: {
            color: "#17b3a3",
          },
        },
        axisLable: {
          interval: 0,
          color: "#333",
        },
      },
      yAxis: [
        {
          type: "value",
          axisLine: {
            lineStyle: {
              color: "#17b3a3",
            },
          },
        },
      ],
      color: ["#2ec7c9", "#b6a2de"],
      series: [
        {
          name: "新增用户",
          data: this.userData.map((item) => item.new),
          type: "bar",
        },
        {
          name: "活跃用户",
          data: this.userData.map((item) => item.active),
          type: "bar",
        },
      ],
    };
    echarts2.setOption(echarts2Option);

    //饼状图
    const echarts3 = echarts.init(this.$refs.echarts3);
    const echarts3Option = {
      tooltip: {
        trigger: "item",
      },
      color: [
        "#0f78f4",
        "#dd536b",
        "#9462e5",
        "#a6a6a6",
        "#e1bb22",
        "#39c362",
        "#3ed1cf",
      ],
      series: [
        {
          data: this.videoData,
          type: "pie",
        },
      ],
    };
    echarts3.setOption(echarts3Option);
  },
};
</script>

<style lang="less" scoped>
.user {
  display: flex;
  align-items: center;
  padding-bottom: 20px;
  margin-bottom: 20px;
  border-bottom: 1px solid #ccc;
  img {
    margin-right: 40px;
    height: 150px;
    width: 150px;
    border-radius: 50%;
  }
  .userInfo {
    .name {
      line-height: 32px;
      font-size: 32px;
      margin-bottom: 10px;
    }
    .access {
      line-height: 32px;
      color: #999999;
    }
  }
}
.login-info {
  color: #999999;
  p {
    margin-left: 0;
    text-align: left;
    line-height: 28px;
    font-size: 14px;
    span {
      color: #666666;
      margin-left: 60px;
    }
  }
}
.num {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  .icon {
    width: 80px;
    height: 80px;
    font-size: 30px;
    line-height: 80px;
    text-align: center;
    color: #fff;
  }
  .detail {
    margin-left: 15px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    .num {
      font-size: 30px;
      margin-bottom: 10px;
      line-height: 30px;
      height: 30px;
    }
    .str {
      font-size: 14px;
      text-align: center;
      color: #999999;
    }
  }
  .el-card {
    width: 32%;
    margin-bottom: 20px;
  }
}
.graph {
  margin-top: 20px;
  display: flex;
  justify-content: space-between;
  .el-card {
    width: 48%;
  }
}
</style>