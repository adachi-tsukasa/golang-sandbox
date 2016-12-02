import { Component,SimpleChange, OnInit, OnChanges, Input } from '@angular/core';

@Component({
  selector: 'right-panel',
  templateUrl: './rightpanel.component.html',
  styleUrls: ['./rightpanel.component.css']
})
export class RightpanelComponent implements OnInit,OnChanges {
  @Input("id")id;
  works: string;
  constructor() {}

  ngOnInit() {
    this.works = "rightpanel"
  }
  ngOnChanges(changes: {[propKey: string]: SimpleChange}) {
    for (let propName in changes) {
      if(!changes[propName].isFirstChange()){
        console.log(changes[propName].previousValue+"から"+changes[propName].currentValue+"に変更されました");
      }
    }
  }
}
