import {Component, OnInit} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {TrayConfig} from './tray-config.model';
import {MessageService} from "primeng/api";
import {TrayConfigService} from "./tray-config.service";

@Component({
    selector: 'app-trayConfig',
    templateUrl: './tray-config.component.html',
    styleUrls: ['./tray-config.component.scss']
})
export class TrayConfigComponent implements OnInit {

    trayConfigs: TrayConfig[];
    trayConfig: TrayConfig;
    trayConfigDialog: boolean;
    submitted: boolean;

    constructor(private trayConfigService: TrayConfigService, private httpClient: HttpClient, private messageService: MessageService) {
    }


    ngOnInit(): void {
        this.trayConfig = {};
        this.loadData();
    }

    loadData() {
        console.log('loading data...');
        this.trayConfigService.read().subscribe(data => this.trayConfigs = data)
    }

    delete(id: number) {
        this.trayConfigService.delete(id).subscribe(() => console.log('deleted'))
        this.loadData();
    }

    edit(trayConfig: TrayConfig) {
        this.trayConfig = {...trayConfig};
        this.trayConfigDialog = true;
    }

    save() {
        this.submitted = true;
        if (this.trayConfig['Id'] === null || this.trayConfig.Id === undefined) {
            this.trayConfigService.create(this.trayConfig)
                .subscribe(value => console.log(JSON.stringify(value)));
            this.messageService.add({
                severity: 'success',
                summary: 'Successful',
                detail: 'TrayConfig created',
                life: 3000
            });
        } else {
            this.trayConfigService.update(this.trayConfig)
                .subscribe(value => console.log(JSON.stringify(value)));
            this.messageService.add({
                severity: 'success',
                summary: 'Successful',
                detail: 'TrayConfig updated',
                life: 3000
            });
        }

        this.loadData();
        this.trayConfigDialog = false;
        this.trayConfig = {};
    }

    openNew() {
        this.trayConfig = {};
        this.submitted = false;
        this.trayConfigDialog = true;
    }

    hideDialog() {
        this.trayConfigDialog = false;
        this.submitted = false;
    }
}
