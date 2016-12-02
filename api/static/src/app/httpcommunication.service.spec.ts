/* tslint:disable:no-unused-variable */

import { TestBed, async, inject } from '@angular/core/testing';
import { HttpcommunicationService } from './httpcommunication.service';

describe('Service: Httpcommunication', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [HttpcommunicationService]
    });
  });

  it('should ...', inject([HttpcommunicationService], (service: HttpcommunicationService) => {
    expect(service).toBeTruthy();
  }));
});
