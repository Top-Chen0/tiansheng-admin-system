/**
 * 网站配置文件
 */

const config = {
  appName: '河南天盛',
  appNamee: '陈俊豪',
  appLogo: 'https://top-chen.com/api/v3/file/source/7032/logo.png?sign=RVlI69MH7t1_JIhskFoTrS9ejS-pw9jQB8Ik_QYDRrU%3D%3A0',
  showViteLogo: true
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    const chalk = require('chalk')
    console.log(
      chalk.green(
        `> 欢迎使用tiansheng-admin-system，开源地址：https://github.com/Top-Chen0/tiansheng-admin-system`
      )
    )
    console.log(
      chalk.green(
        `> 当前版本:v2.5.1 beta`
      )
    )
    console.log(
      chalk.green(
        `> 联系我方式:微信：347517726`
      )
    )
    console.log(
      chalk.green(
        `> https://github.com/Top-Chen0/tiansheng-admin-system`
      )
    )
    console.log(
      chalk.green(
        `> 默认自动化文档地址:http://127.0.0.1:${env.VITE_SERVER_PORT}/swagger/index.html`
      )
    )
    console.log(
      chalk.green(
        `> 默认自动化文档地址:http://127.0.0.1:${env.VITE_SERVER_PORT}/swagger/index.html`
      )
    )
    console.log(
      chalk.green(
        `> 默认前端文件运行地址:http://127.0.0.1:${env.VITE_CLI_PORT}`
      )
    )
    console.log(
      chalk.green(
        `> 如果项目让您获得了收益，希望您能请团队喝杯可乐:微信联系：347517726`
      )
    )
    console.log('\n')
  }
}

export default config
