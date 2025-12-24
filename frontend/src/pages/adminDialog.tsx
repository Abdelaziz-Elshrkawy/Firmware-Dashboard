import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { X } from "lucide-react";
import { useState, type Dispatch, type JSX, type SetStateAction } from "react";
import { ViewKey } from "../types/adminDialog.d";

export default function AdminDialog({
  open,
  setOpen,
}: {
  open: number | null;
  setOpen: Dispatch<SetStateAction<number | null>>;
}) {
  const [activeView, setActiveView] = useState<ViewKey>(ViewKey.Menu);
  const [selectedDevice, setSelectedDevice] = useState<string | null>(null);
  const [selectedFirmware, setSelectedFirmware] = useState<string | null>(null);

  const containerClass =
    "flex flex-col gap-2 p-2 border-2 rounded-md shadow-md";

  const views: Record<ViewKey, JSX.Element> = {
    [ViewKey.Menu]: (
      <div className="flex flex-col gap-2">
        <div className={containerClass}>
          <Button
            className="cursor-pointer btn"
            onClick={() => {
              setActiveView(ViewKey.DeviceList);
            }}
          >
            Devices
          </Button>
          <Button className="cursor-pointer btn">Add Device</Button>
        </div>
        <div className={containerClass}>
          <Button
            className="cursor-pointer btn"
            onClick={() => setActiveView(ViewKey.FirmwaresList)}
          >
            Firmwares
          </Button>
          <Button className="cursor-pointer btn">Add Firmware</Button>
        </div>
      </div>
    ),
    [ViewKey.DeviceList]: (
      <div>
        {["Device 1", "Device 2"].map((d) => (
          <div key={d} className="flex justify-between mb-2">
            <span>{d}</span>
            <Button
              className="cursor-pointer btn btn-sm"
              onClick={() => {
                setSelectedDevice(d);
                setActiveView(ViewKey.DeviceEdit);
              }}
            >
              Edit
            </Button>
          </div>
        ))}
        <Button
          className="mt-2 cursor-pointer btn"
          onClick={() => setActiveView(ViewKey.Menu)}
        >
          Back
        </Button>
      </div>
    ),
    [ViewKey.DeviceEdit]: (
      <div>
        <p>Editing {selectedDevice}</p>
        <Button
          className="mt-2 cursor-pointer btn"
          onClick={() => setActiveView(ViewKey.DeviceList)}
        >
          Back to List
        </Button>
      </div>
    ),
    [ViewKey.FirmwaresList]: (
      <div>
        {["Firmware 1", "Firmware 2"].map((f) => (
          <div key={f} className="flex justify-between mb-2">
            <span>{f}</span>
            <Button
              className="cursor-pointer btn btn-sm"
              onClick={() => {
                setSelectedFirmware(f);
                setActiveView(ViewKey.FirmwareEdit);
              }}
            >
              Edit
            </Button>
          </div>
        ))}
        <Button
          className="mt-2 cursor-pointer btn"
          onClick={() => setActiveView(ViewKey.Menu)}
        >
          Back
        </Button>
      </div>
    ),
    [ViewKey.FirmwareEdit]: (
      <div>
        <p>Editing {selectedFirmware}</p>
        <Button
          className="mt-2 cursor-pointer btn"
          onClick={() => setActiveView(ViewKey.FirmwaresList)}
        >
          Back to List
        </Button>
      </div>
    ),
  };

  return (
    <Dialog open={!!open}>
      <DialogContent className="w-[90%] [&>button]:hidden">
        <DialogHeader>
          <DialogTitle className="flex justify-between">
            {activeView}{" "}
            <X
              onClick={() => {
                setOpen(null);
                setActiveView(ViewKey.Menu);
              }}
              className="font-bold transition-all cursor-pointer hover:text-red-500"
            />
          </DialogTitle>
        </DialogHeader>
        {views[activeView]}
      </DialogContent>
    </Dialog>
  );
}
