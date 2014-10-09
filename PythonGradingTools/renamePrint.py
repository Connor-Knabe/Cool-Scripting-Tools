# Instructions:
# 1) Bulk download labs
# 2) Unzip labs
# 3) Place this script in the folder
# 4) Run 
import os
import shutil

outFile = open('./Grades.txt', 'w+')
pwprtName = ""
pwprtArry = []

#Creates folder for text documents
if (os.path.isdir("TextDocs")):
	print "TextDocs folder already exists"
else:
	os.mkdir("TextDocs")
if (os.path.isdir("Zips")):
	print "Zips folder already exists"
else:
	os.mkdir("Zips")

#Loops through all files changing the name to pwprt.php
#Also adds the pwprt to an array
for filename in os.listdir("."):
	if filename.startswith("Lab ") and "php" in filename:
		pwprtName = filename.split('_')[1]+".php"
		os.rename(filename, pwprtName)
		pwprtArry.append(pwprtName.split(".")[0])
	elif ("sql" in filename):
		pwprtName = filename.split('_')[1]+".sql"
		os.rename(filename, pwprtName)
		pwprtArry.append(pwprtName.split(".")[0])
	elif("txt" in filename and not "Grades.txt" in filename):
		shutil.move(filename,"TextDocs")
	else:		
		shutil.move(filename,"Zips")



#Sorts array and prints to Grades file
pwprtArry.sort()
for j in range (len(pwprtArry)):
    outFile.write(pwprtArry[j] + " -" + "\n")



