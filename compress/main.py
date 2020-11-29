import numpy as np
import csv

binaryArray = []
with open('binary.txt') as f:
	reader = csv.reader(f)
	for line in reader:
		binaryRep = []
		for letter in line[0]:
			binaryRep.append(int(letter))
		binaryArray.append(binaryRep)
data = np.array(binaryArray)
print(data[9])
print(data[99])
print(data[999])
print(data[9999])
print(data[99999])
print(data[999999])
print(data[9999999])
print(data[99999999])
