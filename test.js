// my_script.js

const args = process.argv.slice(2);
const num1 = parseInt(args[0]);
const num2 = parseInt(args[1]);

if (isNaN(num1) || isNaN(num2)) {
    console.error("参数必须是数字");
    process.exit(1); // 错误退出
}

const sum = num1 + num2;
console.log(sum); // 将结果打印到标准输出
process.exit(0);   // 正常退出