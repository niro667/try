Expected = "Hello, Test!"
OUTPUT=$ (node -e "console.log(require('./src/app')('Test'))")

if ["$OUTPUT" == "$EXPECTED"];then

echo = " test pass "

exit 0

else 

echo = "a7a"

exit 1

fi

