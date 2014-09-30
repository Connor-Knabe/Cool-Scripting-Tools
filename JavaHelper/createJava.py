import sys
import os
import subprocess

#to run python createJava.py fileNamehere
if(len(sys.argv)<2):
	print ("Usage: createJava.py fileName")
	sys.exit(0)
fName = sys.argv[1]

def compile():

if(os.path.isfile(fName+".java")):
    print ("Compiling")
    subprocess.call("javac " + fName +".java", shell=True)
    print ("Output below\n")
    subprocess.call("java " + fName +" ", shell=True)
else if (os.path.isfile(fName)):
    print ("Compiling")
    subprocess.call("javac " + fName, shell=True)
    print ("Output below\n")
    subprocess.call("java " + fName +" ", shell=True)
else:
    print ("New file created")
    file = open(fName + ".java", 'w+')
    file.write("public class " + fName + " {\n    public static void main(String[] args) {\n        \n        \n    }\n}")
    


