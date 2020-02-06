import request from '@/utils/request'

const file = {
    uploadImage(data) {
        return request.post('/file/image', data, {
            headers: {
                'Content-Type': 'multipart/form-data'
            },
        })
    },
    imageList(data) {
        return request.get('/file/image', {params: data})
    },
    deleteImage(id) {
        return request.delete('/file/image/'+id)
    }
}

export default file
