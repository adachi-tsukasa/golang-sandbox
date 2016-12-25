import { Response } from '@angular/http';
import { Component,SimpleChange, OnInit, OnChanges, Input } from '@angular/core';
import { FeedService } from '../feed.service';
import { HttpcommunicationService } from '../httpcommunication.service';

@Component({
  selector: 'right-panel',
  providers: [FeedService, HttpcommunicationService],
  templateUrl: './rightpanel.component.html',
  styleUrls: ['./rightpanel.component.css']
})
export class RightpanelComponent implements OnInit,OnChanges {
  @Input("id")id;
  works: string;
  feed: any;
  constructor(private feedservice: FeedService) {}

  ngOnInit() {
    this.works = "rightpanel"
  }
  ngOnChanges(changes: {[propKey: string]: SimpleChange}) {
    let targetValue;
    for (let propName in changes) {
        targetValue = changes[propName].currentValue;
    //   if(!changes[propName].isFirstChange()){
    //     console.log(changes[propName].previousValue+"から"+changes[propName].currentValue+"に変更されました");
    //   }
    }
    if (!targetValue) {
      return;
    }
    this.feedservice.getFeed(targetValue).subscribe((res: Response) =>{
      this.feed = res.json();
    });

  }
}
