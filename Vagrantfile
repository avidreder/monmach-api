# Keep vagrant updated!
Vagrant.require_version ">= 1.8.1"

Vagrant.configure(2) do |config|
  config.vm.provider "virtualbox"
  config.vm.box = "ubuntu/trusty64"

  config.vm.provider "virtualbox" do |v|
    v.memory = 2048
  end

  # Private Network
  config.vm.network :private_network, ip:"192.168.88.8"
  config.ssh.insert_key = false

  # ansible!
  config.vm.provision "ansible" do |ansible|
    ansible.playbook = "ansible/playbook.yml"
    ansible.limit = "localhost"
    ansible.inventory_path = "ansible/inventory"
    ansible.verbose = "v"
    # ansible.ask_sudo_pass = true
  end

  # use `vagrant rsync` or `vagrant rsync-auto` to push into the ONLY if you
  # NEED to. `vagrant provision` will let ansible do it, which is the way a real
  # deployment would work.
  config.vm.synced_folder ".", "/srv/monmach-api", type: "rsync",
    rsync__exclude: ".git/",
    rsync__args: ["--copy-links", "--verbose", "--archive", "--delete", "-z"]
end
