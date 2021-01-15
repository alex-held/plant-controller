import {Component, OnInit} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {MessageService} from "primeng/api";
import {TrayService} from "./tray.service";
import {Tray} from './tray.model';

@Component({
    selector: 'app-tray',
    templateUrl: './tray.component.html',
    styleUrls: ['./tray.component.scss']
})
export class TrayComponent implements OnInit {

    trays: Tray[];
    tray: Tray;
    trayDialog: boolean;
    submitted: boolean;

    constructor(private trayService: TrayService, private httpClient: HttpClient, private messageService: MessageService) {
    }


    ngOnInit(): void {
        this.tray = {};
        this.loadData();
    }


    loadData() {
        console.log('loading data...');
        this.trayService.read().subscribe(data => this.trays = data)
    }

    delete(id: number) {
        this.trayService.delete(id).subscribe(() => console.log('deleted'))
        this.loadData();
    }

    edit(tray: Tray) {
        this.tray = {...tray};
        this.trayDialog = true;
    }

    save() {
        this.submitted = true;
        if (this.tray.Id=== null || this.tray.Id === undefined) {
            this.trayService.create(this.tray)
                .subscribe(value => console.log(JSON.stringify(value)));
            this.messageService.add({
                severity: 'success',
                summary: 'Successful',
                detail: 'Tray created',
                life: 3000
            });
        } else {
            this.trayService.update(this.tray)
                .subscribe(value => console.log(JSON.stringify(value)));
            this.messageService.add({
                severity: 'success',
                summary: 'Successful',
                detail: 'Tray updated',
                life: 3000
            });
        }

        this.loadData();
        this.trayDialog = false;
        this.tray = {};
    }

    openNew() {
        this.tray = {};
        this.submitted = false;
        this.trayDialog = true;
    }

    hideDialog() {
        this.trayDialog = false;
        this.submitted = false;
    }
}
