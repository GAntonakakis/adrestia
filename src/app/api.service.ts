import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

interface Response {
  message: string;
}

interface RequestData {
  firstname: string;
  lastname: string;
  age: string;
  gender: string;
  ethnicity: string;
}

@Injectable({
  providedIn: 'root',
})
export class ApiService {
  apiUrl = 'http://localhost:8080/api/hello';
  dataUrl = 'http://localhost:8080/api/data';

  constructor(private http: HttpClient) {}

  getMessage(): Observable<Response> {
    return this.http.get<Response>(this.apiUrl);
  }

  sendData(data: RequestData): Observable<Response> {
    return this.http.post<Response>(this.dataUrl, data);
  }
}
