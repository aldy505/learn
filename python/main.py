def calculate(typeof, num1, num2):
	return {
		"multiply": num1 * num2,
		"addition": num1 + num2,
		"division": num1 / num2
	}.get(typeof, "Invalid input, please try again")
	# if typeof == "multiply":
	# 	result = num1 * num2
	# elif typeof == "addition":
	# 	result = num1 + num2
	# elif typeof == "division":
	# 	result = num1 / num2
	# else:
	# 	result = "Invalid input. Please try again"
	# return result

print(calculate("multiply", 3, 10))
print(calculate("blah", 10, 10))
print(calculate("addition", 20, 50))