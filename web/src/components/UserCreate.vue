<style>
.createDialog.hostCreate .lbsList.el-col-20 .el-input__inner {
  vertical-align: -1px;
}
</style>
<template>
  <el-dialog top='25vh' :title="dialogTitle" :close-on-click-modal="false" :visible.sync="dialogFormVisible" class="createDialog hostCreate" width='50%' :before-close="doCancelCreate">
    <el-form :model="form" :rules="rules" ref="ruleForm">
      <el-form-item label="用户" prop='user'>
        <el-input v-model.trim="form.user" placeholder="请输入用户" maxlength="64" auto-complete="on" ></el-input>
      </el-form-item>
      <el-form-item label="昵称" prop='name'>
        <el-input v-model.trim="form.name" maxlength="64" auto-complete="off" ></el-input>
      </el-form-item>
      <el-form-item label="Email" prop='email'>
        <el-input v-model.trim="form.email" placeholder="" maxlength="64" auto-complete="off" ></el-input>
      </el-form-item>
      <el-form-item label="WeSite" prop='web_site'>
        <el-input v-model.trim="form.web_site" placeholder="" type="text" maxlength="64" auto-complete="off" ></el-input>
      </el-form-item>
      <el-form-item label="Address" prop='address'>
        <el-input v-model.trim="form.address" placeholder="" type="text" maxlength="64" auto-complete="off"></el-input>
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button @click="doCancelCreate">取消</el-button>
      <el-button type="primary" @click="doSubmit">确定</el-button>
    </div>
  </el-dialog>
</template>
<script>
import { Message } from 'element-ui';
import { addUser, updateUser } from '../api/backend';

const formData = {
  user: '',
  name: '',
  email: '',
  web_site: '',
  address: ''
};
export default {
  props: ['user'],
  data() {
    return {
      // 是否属于编辑状态
      isEdit: false,
      dialogTitle: '创建用户',
      dialogFormVisible: false,
      form: JSON.parse(JSON.stringify(formData)),
      rules: {

      },
    };
  },
  methods: {
    doCreate(flag, item) {
      this.dialogFormVisible = true;
      this.isEdit = flag;
      if (flag) {
        this.dialogTitle = '编辑用户'
        this.form.id = item.id;
        this.form.user = item.user;
        this.form.name = item.name;
        this.form.web_site = item.web_site;
        this.form.address = item.address;
        this.form.email = item.email;
      } else {
        this.dialogTitle = '创建用户'
        this.form = Object.assign({}, formData);
      }
    },
    doCancelCreate() {
      this.$refs.ruleForm.resetFields();
      this.dialogFormVisible = false;
    },
    doSubmit() {
      this.$refs.ruleForm.validate((valid) => {
        if (valid) {
          const successCallBack = (msg) => {
            this.$emit('getlist');
            Message.success(msg);
            this.dialogFormVisible = false;
          };
          if (this.isEdit) {
            updateUser(JSON.stringify(this.form)).then( () => {
                successCallBack('更新成功');
            })
          } else {
            addUser(JSON.stringify(this.form)).then( () => {
                successCallBack('创建成功');
            })
          }
          return false;
        }
      });
    },
  },
};
</script>
