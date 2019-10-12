import { Component, OnInit, Input } from '@angular/core';
import { HTTPLessonService } from '../../httplesson.service';

@Component({
  selector: 'app-scroll',
  templateUrl: './scroll.component.html',
  styleUrls: ['./scroll.component.scss']
})
export class ScrollComponent implements OnInit {
  @Input() username: string;
  punts = [0, 0, 0, 0, 0];

  ITER = [0, 1, 2, 3, 4];
  cards = ["Algorithms", "Complex analysis", "Linear Optimization", "Parametrized Complexity", "Theory"];
  votes = [5,             5,                    0,                    3,                        5];
  it = 0;
  n = 5;

  constructor(private less: HTTPLessonService) { }

  ngOnInit() {
    var _this = this;
    this.less.infoUser(this.username).subscribe(data => {
      _this.n = data.subject.length;
      _this.cards = data.subject;
      _this.votes = data.feedback;
    });
    this.it = 0;
    this.ITER = Array.from(Array(this.n).keys());
    this.punts = Array(this.n);
  }

  expandLess() {
    this.it--;
  }
  expandMore() {
    this.it++;
  }

  pitch(i: number, event: any) {
    this.punts[i] = event.value;
  }

  revote(i: number) {
    this.votes[i] = 0;
    this.punts[i] = 1;
    this.less.sendFeedback(this.username, 0, this.cards[i]).subscribe();
  }

  vote(i: number) {
    this.less.sendFeedback(this.username, this.punts[i], this.cards[i]).subscribe();
    this.votes[i] = this.punts[i];
  }
}
