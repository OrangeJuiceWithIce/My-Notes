# OOP

---

## overload 重载
能够定义多个同名函数，但参数列表必须不同。

### Template<typename T>
```C++
template<typename T>
void swap(const T*a,const T*b)
{
    T temp = *a;
    *a = *b;
    *b = temp;
}
```
## 
array:静态
vector:动态

```C++
vector<int> evens {2,4,6,8,10};
evens.push_back(12);
evens.insert(evens.begin()+6,14);

for(vector<int>::iterator it=evens.begin();it!=evens.end();it++){
    cout<<*it<<" ";
    cout<<endl;
}
```