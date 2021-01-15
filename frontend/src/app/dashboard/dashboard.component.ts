import {Component, OnInit} from '@angular/core';

interface Shortcut {
    uri: string
    name: string

}

@Component({
    selector: 'app-dashboard',
    templateUrl: './dashboard.component.html',
    styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {
    shortcuts: Shortcut[];

    constructor() {
    }

    ngOnInit(): void {
        this.shortcuts = [
            {
                uri: '/plant',
                name: "Plants"
            },
          {
            uri: '/growconfig',
            name: "GrowConfig"
          },
          {
            uri: '/tray',
            name: "Trays"
          },
        ]
    }

}
