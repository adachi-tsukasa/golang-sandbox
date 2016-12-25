import { Http,Request,Response } from '@angular/http';
import { Injectable } from '@angular/core';
import { HttpcommunicationService,Res } from './httpcommunication.service'
import { Observable } from 'rxjs/Observable';

@Injectable()
export class ListService {

  constructor(private httpservice: HttpcommunicationService) { }

  getList(): Observable<Response> {
    return this.httpservice.get("/charList");
  }
}
