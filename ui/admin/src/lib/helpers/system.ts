export const isScrollAtBottom = (element: Element | null = null): boolean => {
    let scrollTop: number, scrollHeight: number, clientHeight: number
    if (element === null) {
        scrollTop = window.scrollY || document.documentElement.scrollTop;
        scrollHeight = document.documentElement.scrollHeight;
        clientHeight = document.documentElement.clientHeight;
    } else {
        scrollTop = element.scrollTop;
        scrollHeight = element.scrollHeight;
        clientHeight = element.clientHeight;
    }
    return scrollTop + clientHeight >= scrollHeight - 5
}

export const debounce = (fn: Function, ms = 300) => {
    let timeoutId: ReturnType<typeof setTimeout>
    return function (this: any, ...args: any[]) {
        clearTimeout(timeoutId)
        timeoutId = setTimeout(() => fn.apply(this, args), ms)
    };
};

// https://decipher.dev/30-seconds-of-typescript/docs/throttle/
export const throttle = (fn: Function, wait: number = 300) => {
    let inThrottle: boolean,
        lastFn: ReturnType<typeof setTimeout>,
        lastTime: number
    return function (this: any) {
        const context = this,
            args = arguments
        if (!inThrottle) {
            fn.apply(context, args)
            lastTime = Date.now()
            inThrottle = true
        } else {
            clearTimeout(lastFn)
            lastFn = setTimeout(() => {
                if (Date.now() - lastTime >= wait) {
                    fn.apply(context, args)
                    lastTime = Date.now()
                }
            }, Math.max(wait - (Date.now() - lastTime), 0))
        }
    };
};