# 🤖 AI 字幕文件翻译

![image](/gosrt.gif)

## 示例
```bash
# 配置ChatGPT KEY 
# ~/.gosrt/config.yaml
gosrt set key sk-duDRTrlEAAYQg3yWF8F64dD5E44xxxxxxxxxxxxxxxxxxx
gosrt set url https://openkey.cloud/v1 # 使用官方api不需要配置

# 字幕翻译指定srt文件
gosrt srt movie.srt

# 字幕翻译指定目录
gosrt srtpath srtdir
```

## 使用
```bash
NAME:
   gosrt.exe - A new cli application

USAGE:
   gosrt.exe [global options] command [command options]

COMMANDS:
   srt      字幕文件
   srtpath  字幕目录
   set      配置
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```