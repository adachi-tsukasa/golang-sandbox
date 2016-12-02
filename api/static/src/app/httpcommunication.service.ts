import { Injectable } from '@angular/core';
import { Http,Request,Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map'

export interface Res {
}

@Injectable()
export class HttpcommunicationService {

  constructor(private http: Http) { }

  get(url: string): Observable<Response> {
    let status: number = 0;
    let body: any = null;
     return this.http.request(new Request({method: "Get",url: url}));
  }

  handleError(res: Response) {
    if (res.status < 200 && res.status > 400) {
      console.error(res.status + "" + res);
    }
  }
}
