import { Injectable } from '@angular/core';

@Injectable()
export class MessageService {

  constructor() { }

  getMessage(): string {
    return "Hello";
  }
}
