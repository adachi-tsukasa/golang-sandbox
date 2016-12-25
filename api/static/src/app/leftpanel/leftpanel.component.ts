import { Http,Request,Response } from '@angular/http';
import { Component, OnInit, EventEmitter, Output } from '@angular/core';

import { ListService } from '../list.service';
import { HttpcommunicationService } from '../httpcommunication.service';

import { Character } from '../model/character';

import { Observable } from 'rxjs/Observable';

@Component({
  providers: [ListService, HttpcommunicationService],
  selector: 'left-panel',
  templateUrl: './leftpanel.component.html',
  styleUrls: ['./leftpanel.component.css']
})
export class LeftpanelComponent implements OnInit {
  @Output() emitter = new EventEmitter<string>();
  charsList: Character[];
  id: string;

  constructor(private listservice: ListService) {
    // this.emitter = new EventEmitter();
  }


  ngOnInit() {
    this.listservice.getList().subscribe((res: Response) =>{
      this.charsList = res.json();
    });
  }
  selectCharacter(id: string) {
    this.emitter.emit(id);
  }
}
