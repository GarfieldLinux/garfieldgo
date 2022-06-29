
//格式化输出struct:
 bs,_ := json.Marshal(cloud_domain_jd.BackSourceType)
 var out bytes.Buffer
 json.Indent(&out,bs,"","\t")
(":",out.String())
