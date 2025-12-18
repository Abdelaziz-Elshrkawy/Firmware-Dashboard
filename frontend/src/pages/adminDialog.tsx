import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { X } from "lucide-react";
import { useState, type Dispatch, type JSX, type SetStateAction } from "react";

type ViewKey =
  | "menu"
  | "devicesList"
  | "deviceEdit"
  | "firmwaresList"
  | "firmwareEdit";

export default function AdminDialog({
  open,
  setOpen,
}: {
  open: number | null;
  setOpen: Dispatch<SetStateAction<number | null>>;
}) {
  const [activeView, setActiveView] = useState<ViewKey>("menu");
  const [selectedDevice, setSelectedDevice] = useState<string | null>(null);
  const [selectedFirmware, setSelectedFirmware] = useState<string | null>(null);

  const containerClass =
    "flex flex-col gap-2 p-2 border-2 rounded-md shadow-md";

  const views: Record<ViewKey, JSX.Element> = {
    menu: (
      <div className="flex flex-col gap-2">
        <div className={containerClass}>
          <Button
            className="cursor-pointer btn"
            onClick={() => {
              setActiveView("devicesList");
            }}
          >
            Devices
          </Button>
          <Button className="cursor-pointer btn">Add Device</Button>
        </div>
        <div className={containerClass}>
          <Button
            className="cursor-pointer btn"
            onClick={() => setActiveView("firmwaresList")}
          >
            Firmwares
          </Button>
          <Button className="cursor-pointer btn">Add Firmware</Button>
        </div>
      </div>
    ),
    devicesList: (
      <div>
        {["Device 1", "Device 2"].map((d) => (
          <div key={d} className="flex justify-between mb-2">
            <span>{d}</span>
            <Button
              className="cursor-pointer btn btn-sm"
              onClick={() => {
                setSelectedDevice(d);
                setActiveView("deviceEdit");
              }}
            >
              Edit
            </Button>
          </div>
        ))}
        <Button
          className="mt-2 cursor-pointer btn"
          onClick={() => setActiveView("menu")}
        >
          Back
        </Button>
      </div>
    ),
    deviceEdit: (
      <div>
        <p>Editing {selectedDevice}</p>
        <Button
          className="mt-2 cursor-pointer btn"
          onClick={() => setActiveView("devicesList")}
        >
          Back to List
        </Button>
      </div>
    ),
    firmwaresList: (
      <div>
        {["Firmware 1", "Firmware 2"].map((f) => (
          <div key={f} className="flex justify-between mb-2">
            <span>{f}</span>
            <Button
              className="cursor-pointer btn btn-sm"
              onClick={() => {
                setSelectedFirmware(f);
                setActiveView("firmwareEdit");
              }}
            >
              Edit
            </Button>
          </div>
        ))}
        <Button
          className="mt-2 cursor-pointer btn"
          onClick={() => setActiveView("menu")}
        >
          Back
        </Button>
      </div>
    ),
    firmwareEdit: (
      <div>
        <p>Editing {selectedFirmware}</p>
        <Button
          className="mt-2 cursor-pointer btn"
          onClick={() => setActiveView("firmwaresList")}
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
                setActiveView("menu");
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
