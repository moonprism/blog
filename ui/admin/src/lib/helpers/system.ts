
export function scrollAtBottom(): boolean {
    const scrollTop = window.scrollY
    const scrollHeight = document.documentElement.scrollHeight
    const clientHeight = window.innerHeight
    return scrollTop + clientHeight >= scrollHeight - 5
}