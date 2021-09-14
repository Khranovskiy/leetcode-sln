//https://medium.com/javascript-in-plain-english/simple-things-gone-wrong-how-to-find-out-the-number-of-digits-in-a-javascript-number-c0a78c6acbec
function getDigits(value) {
    const number = Number(value);
    const string = number.toString();
    const hasExponent = string.indexOf('e') !== -1; // is the number written using the scientific notation

    if (hasExponent) {
        const exponent = string.split('e')[1].substring(1);
        return parseInt(exponent) + 1;
    }

    return number.toString().length;
}

