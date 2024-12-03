def iterate(array):
    if array[len(array)-1]==array[0]:
        return false
    if int(array[len(array)-1])<int(array[0]):
        constant=-3
    else:
        constant=3
    for i in range(0,len(array)-1):
        diff = int(array[i])-int(array[i+1])
        if abs(diff)>abs(constant):
            return False #diff is too big
        if (diff>0):
            if (constant>0):
                return False
        if (diff<0):
            if (constant<0):
                return False
        if (diff == 0):
            return False #diff is zero
    return True

count=0
count2=0
f=open("input.txt","r")
for i in range(0,1000):
    a=f.readline()
    a=a.split(" ")
    #check if monotonic
    if iterate(a):
        count+=1
f.close()
 
print(count)
