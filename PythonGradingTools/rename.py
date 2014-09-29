import os
for filename in os.listdir("."):
	if filename.startswith("Lab ") and "php" in filename:
		os.rename(filename, filename.split('_')[1]+".php")
