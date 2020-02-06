export default {
    namespaced: true,
    state: {
        statusList: [
            {
                status: 1,
                text: 'draft',
                type: 'warning'
            },
            {
                status: 2,
                text: 'publish',
                type: 'success'
            },
            {
                status: 3,
                text: 'system',
                type: 'info'
            }
        ]
    },
    getters: {
        statusList: state => state.statusList
    }
}