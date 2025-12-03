import http from 'k6/http';
import {sleep} from 'k6';

export const options = {
    duration: '60s',
    vus: 50, 
};

export default function(){
    http.get('http://localhost:8080/fib?n=40');
    sleep(0.1);
}