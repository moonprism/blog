export function insertStr(str, start, newStr) {
    return str.slice(0, start) + newStr + str.slice(start)
}