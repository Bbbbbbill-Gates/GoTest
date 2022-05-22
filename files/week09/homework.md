# Q1 解决粘包方式
1.固定长度，每个包长度固定，长度不够说明包未完全收到，长度大于固定长度说明大于一个包
2.包分隔符，根绝固定界定符确定包的边界，但是需要考虑分隔符转义的问题
3.包携带长度信息，每个包携带自己的长度信息，比固定长度灵活

# Q2 socket connection 中解码出 goim 协议

``` go

func convertByteOrder(string) (int) {
    // ...
}

func decoderGoim(goimMessage string) (err int) {
  if string.size() < 4 
      return -1;
  len = convertByteOrder(goimMessage[0:3])
  if string.size() < len
      return -1;
  headerLen = convertByteOrder(goimMessage[4:5])
  ver = convertByteOrder(goimMessage[6:7])
  opt = goimMessage[7:8]
  seq = convertByteOrder(goimMessage[9:12])
  body = convertByteOrder(goimMessage[13:len - 1])
  return 0;
}

```
