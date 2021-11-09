# Copyright (c) 2020 Gitpod GmbH. All rights reserved.
# Licensed under the GNU Affero General Public License (AGPL).
# See License-AGPL.txt in the project root for license information.

FROM scratch

# 注意：这必须是图像中的第一层，s.t. 那个 blobserve
# 可以服务IDE主机。 即使在这条线之前移动 WORKDIR 也会破坏事情。
COPY components-supervisor-frontend--app/node_modules/@gitpod/supervisor-frontend/dist/ /.supervisor/frontend/

WORKDIR "/.supervisor"
COPY components-supervisor--app/supervisor \
     supervisor-config.json \
     components-workspacekit--app/workspacekit \
     components-workspacekit--fuse-overlayfs/fuse-overlayfs \
     components-gitpod-cli--app/gitpod-cli \
     ./
WORKDIR "/.supervisor/dropbear"
COPY components-supervisor--dropbear/dropbear \
     components-supervisor--dropbear/dropbearkey \
     ./

ENTRYPOINT ["/.supervisor/supervisor"]