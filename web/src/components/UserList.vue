<style>
.page-content {
  margin: 20px;
}
</style>

<template>
<div class="page-content">
  <el-button type="success" plain @click="getList()">刷新</el-button>
  <el-button type="primary" plain  @click="$refs.create.doCreate(false)">创建用户</el-button>
  <el-table
    :data="tableData"
    style="width: 100%">
    <el-table-column prop="user" label="User"></el-table-column>
    <el-table-column prop="name" label="NickName" ></el-table-column>
    <el-table-column prop="email" label="Email" />
    <el-table-column prop="web_site" label="WebSite" />
    <el-table-column prop="address" label="Address"/>
    <el-table-column fixed="right" label="Operations" width="150">
      <template slot-scope="scope">
        <el-button @click="$refs.create.doCreate(true, scope.row)"  type="text" size="small">编辑</el-button>
        <el-button @click="handleDelete(scope.row.id)"  type="text" size="small">删除</el-button>
      </template>
    </el-table-column>
  </el-table>
  <user-create ref="create" v-on:getlist="getList"></user-create>
</div>

</template>
<script>
import UserCreate from './UserCreate.vue'
import { UserList,deleteUser } from '../api/backend';
  export default {
    name: "UserList",
    components: {
     UserCreate
    },
    created() {
      this.getList()
    },
    methods: {
      handleDelete(id) {
        deleteUser(id).then(
          response => {
            console.log(response)
            this.getList()
          }
        )

      },
      getList() {
        UserList().then(
          response => {
            this.tableData = response.data
          }
        )
      },
    },
    data() {
      return {
        tableData: []
      }
    }
  }
</script>