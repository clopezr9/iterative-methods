import numpy as np 

def isDDM(m, n) :
 
    # for each row
    for i in range(0, n) :        
     
        # for each column, finding
        # sum of each row.
        sum = 0
        for j in range(0, n) :
            sum = sum + abs(m[i][j])    
 
        # removing the
        # diagonal element.
        sum = sum - abs(m[i][i])
 
        # checking if diagonal
        # element is less than
        # sum of non-diagonal
        # element.
        if (abs(m[i][i]) < sum) :
            m[i][i] += sum 
            return False
    return True

# def randomGenerator() :
#     matrix = np.random.rand(501,501)
#     return matrix

# Driver Code
#  50 300 700 1000 1500 2000

n = 1500
m = np.random.rand(1500,1501)

while True :
    if((isDDM(m, n))) :
        print (m)
        np.savetxt("1500-4.csv", m, 
              delimiter = ",")
        break
    else :
        continue


 
