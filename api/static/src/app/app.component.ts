import { Component} from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  targetId: string;
  title = 'app works!';

  receiveEvent(id) {
    console.log("AppComponent:"+id);
    // TODO: rightpanelで必要な物を作る

    // to rightpanel
    this.targetId = id;
  }
}