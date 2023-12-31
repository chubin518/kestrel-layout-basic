import axios from 'axios';
import type { AxiosRequestConfig, AxiosResponse } from 'axios';
import { Message, Modal } from '@arco-design/web-vue';

export interface HttpResponse<T = unknown> {
  status: number;
  message: string;
  data: T;
}

if (import.meta.env.VITE_API_BASE_URL) {
  axios.defaults.baseURL = import.meta.env.VITE_API_BASE_URL;
}

axios.interceptors.request.use(
  (config: AxiosRequestConfig) => {
    // let each request carry token
    // this example using the JWT token
    // Authorization is a custom headers key
    // please modify it according to the actual situation

    return config;
  },
  (error) => {
    // do something
    return Promise.reject(error);
  }
);

// add response interceptors
axios.interceptors.response.use(
  (response: AxiosResponse<HttpResponse>) => {
    const res = response.data;
    if (res.status !== 200) {
      Message.error({
        content: res.message || 'Error',
        duration: 5 * 1000,
      });
      if (res.status === 401) {
        Modal.error({
          title: 'Confirm logout',
          content:
            'You have been logged out, you can cancel to stay on this page, or log in again',
          okText: 'Re-Login',
          async onOk() {
            // todo logout
            window.location.reload();
          },
        });
      }
      return Promise.reject(new Error(res.message || 'Error'));
    }
    return res;
  },
  (error) => {
    Message.error({
      content: error.msg || 'Request Error',
      duration: 5 * 1000,
    });
    return Promise.reject(error);
  }
);
