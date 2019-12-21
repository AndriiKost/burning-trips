import axios from 'axios'
import { AWS_LAMBDA_GETSIGNEDURL_ENDPOINT } from '../config/index';

export default {
  getSignedURL (file) {
    let endpoint = AWS_LAMBDA_GETSIGNEDURL_ENDPOINT
    let payload = {
      filePath: file.name,
      contentType: file.type
    }
    return axios.post(endpoint, payload)
      .then((res) => {
        return Promise.resolve(res.data.url || '/')
      })
      .catch((err) => {
        console.error(err)
        return Promise.reject('/')
      })
  }
}