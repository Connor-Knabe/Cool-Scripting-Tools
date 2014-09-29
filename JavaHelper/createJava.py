import sys
import os
import subprocess

#to run python createJava.py fileNamehere

fName = sys.argv[1]
if(os.path.isfile(fName+".java")):
	
    subprocess.call("javac " + fName +".java", shell=True)
    subprocess.call("java " + fName +" ", shell=True)
else:
    print ("New file created")
    file = open(fName + ".java", 'w+')
    file.write("public class " + fName + " {\n    public static void main(String[] args) {\n        \n        \n    }\n}")
    


