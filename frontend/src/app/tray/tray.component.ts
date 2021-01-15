import {Component, OnInit} from '@angular/core';

interface Tray {
    id?: number
    status: string
    plant?: string
    category?: string
}

@Component({
    selector: 'app-tray',
    templateUrl: './tray.component.html',
    styleUrls: ['./tray.component.scss']
})
export class TrayComponent implements OnInit {
    trays: Tray[];

    constructor() {
    }

    ngOnInit(): void {
        this.trays = [
            {
                id: 1,
                plant: 'Alfalfa',
                category: "tray",
                status: 'active'
            },
            {
                id: 2,
                plant: '',
                category: "tray",
                status: 'empty'
            },
            {
                id: 3,
                plant: 'Rucola',
                category: "tray",
                status: 'empty'
            },
            {
                id: 4,
                plant: 'Basilikum',
                category: "pot",
                status: 'active'
            },
        ];
    }
}
