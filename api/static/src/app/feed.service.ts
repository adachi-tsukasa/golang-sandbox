import { Http,Request,Response } from '@angular/http';
import { Injectable } from '@angular/core';
import { HttpcommunicationService,Res } from './httpcommunication.service'
import { Observable } from 'rxjs/Observable';

@Injectable()
export class FeedService {

  constructor(private httpservice: HttpcommunicationService) { }

  getFeed(targetId: string): Observable<Response> {
    // this.arrangeMentPosts(this.httpservice.get("/feed/" + targetId), '5');

    return this.httpservice.get("/feed/" + targetId);
    
  }

  arrangeMentPosts(rawPosts: Observable<Response>, limitDate: string) {
    let retunablePosts;
    rawPosts.subscribe(function(posts: Response) {

    });
  }
}
