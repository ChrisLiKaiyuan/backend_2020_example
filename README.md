# backend_2020_example
learning gin & gorm.

## methods
> 增： <br>
```
POST /student
{
	staffID: "",
	staffName: "",
	phone: ""
}
```
> 删： <br>
```
DELETE /student?id=
```
> 改： <br>
```
PUT /student?id=
{
	staffName: "",
	phone: ""
}
```
> 查： <br>
```
GET /student
GET /student?id=
```
查那里，最开始的接口是`/student`，是查询所有人的信息
<br>
后面改成`/student?id=`，按照学号查找，但是旧的也没删掉

## TODO
- [ ] validator