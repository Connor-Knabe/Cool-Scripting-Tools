import os
arr = []
outFile = open('./names.txt', 'w+')
for filenames in os.listdir('.'):
    arr.append(filenames.split(".")[0])
arr.sort()
for j in range (len(arr)):
    outFile.write(arr[j] + " -" + "\n")
